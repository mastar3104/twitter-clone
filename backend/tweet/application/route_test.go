package application

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/application/usecase"
	"github.com/mastar3104/twitter-clone/tweet/application/view"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
 * ツイート系APIのインテグレーションテストを実装
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

	t.Run("GET /v1/tweets/:userId", func(t *testing.T) {
		tweetID1 := value.CreateTweetID()
		tweetID2 := value.CreateTweetID()
		content1 := "ツイート1"
		content2 := "ツイート2"
		testUserID := uuid.New().String()
		testUserName := "ツイートテストユーザ"
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUserID, testUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", tweetID1.ToString(), testUserID, content1, "2024-01-01 00:00:01", "2024-01-01 00:00:01")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", tweetID2.ToString(), testUserID, content2, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		// SELECT対象外(別ユーザ)のツイート
		noiseUserID := uuid.New().String()
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", noiseUserID, "ノイズユーザ", "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", value.CreateTweetID().ToString(), noiseUserID, "noise", "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		req, _ := http.NewRequest("GET", testServer.URL+"/v1/tweets/"+testUserID, nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)
		resBody, _ := io.ReadAll(res.Body)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("指定したユーザIDのツイート情報が全てレスポンスされること", func(t *testing.T) {
			var actual view.Tweets
			if err := json.Unmarshal(resBody, &actual); err != nil {
				t.Fatalf(err.Error())
			}
			t.Run("jsonサイズがINSERTしたレコード件数と一致すること", func(t *testing.T) {
				if len(actual.Tweets) != 2 {
					t.Fatalf("tweets size unmatched. actual: %d, expected: %d", len(actual.Tweets), 2)
				}
			})
			t.Run("1件目のユーザIDが"+testUserID+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].UserId != testUserID {
					t.Fatalf("tweets[0].userId　unmatched. actual: %s, expected: %s", actual.Tweets[0].UserId, testUserID)
				}
			})
			t.Run("1件目のユーザ名が"+testUserName+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].UserName != testUserName {
					t.Fatalf("tweets[0].userName　unmatched. actual: %s, expected: %s", actual.Tweets[0].UserName, testUserName)
				}
			})
			t.Run("1件目のツイートIDが"+tweetID1.ToString()+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].TweetId != tweetID1.ToString() {
					t.Fatalf("tweets[0].tweetId　unmatched. actual: %s, expected: %s", actual.Tweets[0].TweetId, tweetID1.ToString())
				}
			})
			t.Run("1件目のツイート内容が"+content1+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].Content != content1 {
					t.Fatalf("tweets[0].content　unmatched. actual: %s, expected: %s", actual.Tweets[0].Content, content1)
				}
			})
			t.Run("1件目のツイート時刻が 2024-01-01 00:00:01(UTC)のUNIXTIME と一致すること", func(t *testing.T) {
				if actual.Tweets[0].TweetTime != 1704067201 {
					t.Fatalf("tweets[0].tweetTime　unmatched. actual: %d, expected: %d", actual.Tweets[0].TweetTime, 1704034800)
				}
			})
		})
	})

	t.Run("GET /v1/timeline/:userId", func(t *testing.T) {
		tweetID1 := value.CreateTweetID()
		tweetID2 := value.CreateTweetID()
		content1 := "ツイート1"
		content2 := "ツイート2"
		testUserID := uuid.New().String()
		testUserName := "ツイートテストユーザ"
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUserID, testUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", tweetID1.ToString(), testUserID, content1, "2024-01-01 00:00:01", "2024-01-01 00:00:01")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", tweetID2.ToString(), testUserID, content2, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO timeline(tweet_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", tweetID1.ToString(), testUserID, "2024-01-01 00:00:10", "2024-01-01 00:00:10")
		handler.Exec("INSERT INTO timeline(tweet_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", tweetID2.ToString(), testUserID, "2024-01-01 00:00:15", "2024-01-01 00:00:15")
		// SELECT対象外(別ユーザ)のツイートとタイムライン
		noiseUserID := uuid.New().String()
		noiseTweetID := value.CreateTweetID().ToString()
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", noiseUserID, "ノイズユーザ", "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", noiseTweetID, noiseUserID, "noise", "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO timeline(tweet_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", noiseTweetID, noiseUserID, "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		req, _ := http.NewRequest("GET", testServer.URL+"/v1/timeline/"+testUserID, nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)
		resBody, _ := io.ReadAll(res.Body)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("指定したユーザIDのツイート情報が全てレスポンスされること", func(t *testing.T) {
			var actual view.Tweets
			if err := json.Unmarshal(resBody, &actual); err != nil {
				t.Fatalf(err.Error())
			}
			t.Run("jsonサイズがINSERTしたレコード件数と一致すること", func(t *testing.T) {
				if len(actual.Tweets) != 2 {
					t.Fatalf("tweets size unmatched. actual: %d, expected: %d", len(actual.Tweets), 2)
				}
			})
			t.Run("1件目のユーザIDが"+testUserID+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].UserId != testUserID {
					t.Fatalf("tweets[0].userId　unmatched. actual: %s, expected: %s", actual.Tweets[0].UserId, testUserID)
				}
			})
			t.Run("1件目のユーザ名が"+testUserName+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].UserName != testUserName {
					t.Fatalf("tweets[0].userName　unmatched. actual: %s, expected: %s", actual.Tweets[0].UserName, testUserName)
				}
			})
			t.Run("1件目のツイートIDが"+tweetID1.ToString()+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].TweetId != tweetID1.ToString() {
					t.Fatalf("tweets[0].tweetId　unmatched. actual: %s, expected: %s", actual.Tweets[0].TweetId, tweetID1.ToString())
				}
			})
			t.Run("1件目のツイート内容が"+content1+"と一致すること", func(t *testing.T) {
				if actual.Tweets[0].Content != content1 {
					t.Fatalf("tweets[0].content　unmatched. actual: %s, expected: %s", actual.Tweets[0].Content, content1)
				}
			})
			t.Run("1件目のツイート時刻が 2024-01-01 00:00:01(UTC)のUNIXTIME と一致すること", func(t *testing.T) {
				if actual.Tweets[0].TweetTime != 1704067201 {
					t.Fatalf("tweets[0].tweetTime　unmatched. actual: %d, expected: %d", actual.Tweets[0].TweetTime, 1704034800)
				}
			})
		})
	})

	t.Run("POST /v1/tweets/:userId", func(t *testing.T) {
		testUserID := uuid.New().String()
		testUserName := "ツイートテストユーザ"
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testUserID, testUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		testContent := "テストツイート"
		payload := usecase.TweetPostBody{Content: testContent}
		jsonValue, _ := json.Marshal(payload)
		res, _ := http.Post(testServer.URL+"/v1/tweets/"+testUserID, "application/json", bytes.NewBuffer(jsonValue))
		resBody, _ := io.ReadAll(res.Body)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})
		var actual view.Tweet
		t.Run("投稿されたツイート情報がレスポンスされること", func(t *testing.T) {
			if err := json.Unmarshal(resBody, &actual); err != nil {
				t.Fatalf(err.Error())
			}
			t.Run("レスポンスされたユーザIDが投稿したユーザIDと一致すること", func(t *testing.T) {
				if actual.UserId != testUserID {
					t.Fatalf("userId　unmatched. actual: %s, expected: %s", actual.UserId, testUserID)
				}
			})
			t.Run("レスポンスされたユーザ名が投稿したユーザ名と一致すること", func(t *testing.T) {
				if actual.UserName != testUserName {
					t.Fatalf("userName　unmatched. actual: %s, expected: %s", actual.UserName, testUserName)
				}
			})
			t.Run("レスポンスされたユーザ名が投稿したユーザ名と一致すること", func(t *testing.T) {
				if actual.Content != testContent {
					t.Fatalf("content　unmatched. actual: %s, expected: %s", actual.Content, testContent)
				}
			})
			t.Run("投稿したツイートIDがレスポンスされること", func(t *testing.T) {
				if actual.TweetId == "" {
					t.Fatalf("tweetId Fail.")
				}
			})
		})

		t.Run("ツイート情報がデータベース上に登録されていること", func(t *testing.T) {
			rows := handler.Query("SELECT content FROM tweet WHERE tweet_id = ?", actual.TweetId)
			defer func(rows *sql.Rows) {
				if err := rows.Close(); err != nil {
					t.Fatalf(err.Error())
				}
			}(rows)
			rows.Next()
			var actualContent string
			if err := rows.Scan(&actualContent); err != nil {
				t.Fatalf(err.Error())
			}
			if actualContent != testContent {
				t.Fatalf("insert faild. actual: %s, expected: %s", actualContent, testContent)
			}
		})
	})
}
