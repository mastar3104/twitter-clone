package application

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/user/application/usecase"
	"github.com/mastar3104/twitter-clone/user/application/view"
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/**
 * ユーザ系APIのインテグレーションテストを実装
 */

func TestRoute(t *testing.T) {
	if err := godotenv.Load("../../.env.testing"); err != nil {
		panic(err.Error())
	}
	handler := adapter.GetDatabaseHandler()
	defer handler.Close()
	server := echo.New()
	Route(server, handler)
	testServer := httptest.NewServer(server)

	t.Run("GET /v1/users/:userId", func(t *testing.T) {
		testUserID := value.CreateNewUserID().ToString()
		testUserName := "テストユーザ1"
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUserID, testUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		req, _ := http.NewRequest("GET", testServer.URL+"/v1/users/"+testUserID, nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)
		resBody, _ := io.ReadAll(res.Body)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("指定したIDのユーザ情報がレスポンスされること", func(t *testing.T) {
			var actual view.User
			if err := json.Unmarshal(resBody, &actual); err != nil {
				t.Fatalf(err.Error())
			}
			if actual.UserName != testUserName {
				t.Fatalf("Body Fail.\nResponseBody: %s,\nexpected: %s", actual.UserName, testUserName)
			}
			if actual.UserID != testUserID {
				t.Fatalf("Body Fail.\nResponseBody: %s,\nexpected: %s", actual.UserID, testUserID)
			}
		})
	})

	t.Run("POST /v1/users", func(t *testing.T) {
		t.Run("正常系リクエスト", func(t *testing.T) {
			expectedUserName := "テストユーザ2"
			payload := usecase.UserPostBody{UserName: expectedUserName}
			jsonValue, _ := json.Marshal(payload)
			res, _ := http.Post(testServer.URL+"/v1/users", "application/json", bytes.NewBuffer(jsonValue))
			resBody, _ := io.ReadAll(res.Body)

			t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
				if res.StatusCode != 200 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
				}
			})

			var actual view.User
			t.Run("登録されたユーザ情報がレスポンスされること", func(t *testing.T) {
				if err := json.Unmarshal(resBody, &actual); err != nil {
					t.Fatalf(err.Error())
				}
				if actual.UserName != expectedUserName {
					t.Fatalf("Body Fail.\nResponseBody: %s,\nexpected: %s", actual.UserName, expectedUserName)
				}
				if actual.UserID == "" {
					t.Fatalf("Body Fail.\nResponseBody: %s", actual.UserID)
				}
			})

			t.Run("ユーザ情報がデータベース上に登録されていること", func(t *testing.T) {
				rows := handler.Query("SELECT user_name FROM user WHERE user_id = ?", actual.UserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				rows.Next()
				var actualUserName string
				if err := rows.Scan(&actualUserName); err != nil {
					t.Fatalf(err.Error())
				}
				if actualUserName != expectedUserName {
					t.Fatalf("insert faild. actual: %s, expected: %s", actualUserName, expectedUserName)
				}
			})
		})

		t.Run("準正常系リクエスト", func(t *testing.T) {
			t.Run("ボディが空である場合、レスポンスステータスコードが400であること", func(t *testing.T) {
				payload := usecase.UserPostBody{}
				jsonValue, _ := json.Marshal(payload)
				res, _ := http.Post(testServer.URL+"/v1/users", "application/json", bytes.NewBuffer(jsonValue))
				if res.StatusCode != 400 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 400)
				}
			})

			t.Run("ボディがJson要素ではない場合、レスポンスステータスコードが400であること", func(t *testing.T) {
				res, _ := http.PostForm(testServer.URL+"/v1/users", url.Values{
					"UserName": {"テストユーザ１"},
				})
				if res.StatusCode != 400 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 400)
				}
			})
		})
	})

	t.Run("POST /v1/user/:userId/follow", func(t *testing.T) {
		t.Run("正常系リクエスト", func(t *testing.T) {
			followerUserID := value.CreateNewUserID().ToString()
			followerUserName := "フォロワー"
			followUserID := value.CreateNewUserID().ToString()
			followUserName := "フォロー"
			// モックデータのINSERT
			handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followerUserID, followerUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
			handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followUserID, followUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")

			payload := usecase.FollowPostBody{FollowUserID: followUserID}
			jsonValue, _ := json.Marshal(payload)
			res, _ := http.Post(testServer.URL+"/v1/users/"+followerUserID+"/follow", "application/json", bytes.NewBuffer(jsonValue))

			t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
				if res.StatusCode != 200 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
				}
			})

			t.Run("ユーザ情報がデータベース上に登録されていること", func(t *testing.T) {
				rows := handler.Query("SELECT follow_user_id FROM follow WHERE user_id = ?", followerUserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				rows.Next()
				var actualFollowUserID string
				if err := rows.Scan(&actualFollowUserID); err != nil {
					t.Fatalf(err.Error())
				}
				if actualFollowUserID != followUserID {
					t.Fatalf("insert faild. actual: %s, expected: %s", actualFollowUserID, followUserID)
				}
			})
		})

		t.Run("準正常系リクエスト", func(t *testing.T) {

			t.Run("リクエストパスに指定したユーザIDが存在しない場合、レスポンスステータスコードが404であること", func(t *testing.T) {
				followerUserID := value.CreateNewUserID().ToString()
				followUserID := value.CreateNewUserID().ToString()
				followUserName := "フォロー"
				// モックデータのINSERT
				// followerのINSERTを実行しない handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followerUserID, followerUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
				handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followUserID, followUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
				payload := usecase.FollowPostBody{FollowUserID: followUserID}
				jsonValue, _ := json.Marshal(payload)
				res, _ := http.Post(testServer.URL+"/v1/users/"+followerUserID+"/follow", "application/json", bytes.NewBuffer(jsonValue))
				if res.StatusCode != 404 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 400)
				}
			})

			t.Run("リクエストボディに指定したフォローユーザIDが存在しない場合、レスポンスステータスコードが404であること", func(t *testing.T) {
				followerUserID := value.CreateNewUserID().ToString()
				followerUserName := "フォロワー"
				followUserID := value.CreateNewUserID().ToString()
				// モックデータのINSERT
				handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followerUserID, followerUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
				// followのINSERTを実行しない  handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", followUserID, followUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
				payload := usecase.FollowPostBody{FollowUserID: followUserID}
				jsonValue, _ := json.Marshal(payload)
				res, _ := http.Post(testServer.URL+"/v1/users/"+followerUserID+"/follow", "application/json", bytes.NewBuffer(jsonValue))
				if res.StatusCode != 404 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 400)
				}
			})

			t.Run("ボディが空である場合、レスポンスステータスコードが400であること", func(t *testing.T) {
				followerUserID := value.CreateNewUserID().ToString()
				payload := usecase.FollowPostBody{}
				jsonValue, _ := json.Marshal(payload)
				res, _ := http.Post(testServer.URL+"/v1/users/"+followerUserID+"/follow", "application/json", bytes.NewBuffer(jsonValue))
				if res.StatusCode != 400 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
				}
			})

			t.Run("ボディがJson要素ではない場合、レスポンスステータスコードが400であること", func(t *testing.T) {
				followerUserID := value.CreateNewUserID().ToString()
				followUserID := value.CreateNewUserID().ToString()
				res, _ := http.PostForm(testServer.URL+"/v1/users/"+followerUserID+"/follow", url.Values{
					"FollowUserID": {followUserID},
				})
				if res.StatusCode != 400 {
					t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
				}
			})
		})
	})

	t.Run("GET /v1/users", func(t *testing.T) {
		testUser1ID := value.CreateNewUserID().ToString()
		testUser1Name := "テストユーザ1"
		testUser2ID := value.CreateNewUserID().ToString()
		testUser2Name := "テストユーザ2"
		// モックデータのINSERT
		handler.Exec("DELETE FROM timeline")
		handler.Exec("DELETE FROM tweet")
		handler.Exec("DELETE FROM follow")
		handler.Exec("DELETE FROM user")
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUser1ID, testUser1Name, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUser2ID, testUser2Name, "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		req, _ := http.NewRequest("GET", testServer.URL+"/v1/users", nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)
		resBody, _ := io.ReadAll(res.Body)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("全てユーザ情報がレスポンスされること", func(t *testing.T) {
			var actual view.Users
			if err := json.Unmarshal(resBody, &actual); err != nil {
				t.Fatalf(err.Error())
			}
			if len(actual.Users) != 2 {
				t.Fatalf("rows count unmatched. actual: %d,expected: %d", len(actual.Users), 2)
			}
		})
	})
}
