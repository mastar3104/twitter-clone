package application

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	if err := godotenv.Load("../../.env.testing"); err != nil {
		panic(err.Error())
	}
	handler := adapter.GetDatabaseHandler()
	defer handler.Close()
	server := echo.New()
	Route(server, handler)
	testServer := httptest.NewServer(server)

	t.Run("POST /v1/users/:userId/tweets/:tweetId", func(t *testing.T) {
		testTweetUserID := uuid.New().String()
		testTweetUserName := "ツイートユーザ"
		testTweetID := uuid.New().String()
		testFollower1UserID := uuid.New().String()
		testFollower1UserName := "フォロワーユーザ1"
		testFollower2UserID := uuid.New().String()
		testFollower2UserName := "フォロワーユーザ2"
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testTweetUserID, testTweetUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollower1UserID, testFollower1UserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollower2UserID, testFollower2UserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO follow(user_id, follow_user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollower1UserID, testTweetUserID, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO follow(user_id, follow_user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollower2UserID, testTweetUserID, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", testTweetID, testTweetUserID, "テストツイート１", "2024-01-01 00:00:01", "2024-01-01 00:00:01")
		// タイムライン生成対象外のレコード
		noiseUserID := uuid.New().String()
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", noiseUserID, "ノイズユーザ", "2024-01-01 00:00:00", "2024-01-01 00:00:00")

		req, _ := http.NewRequest("POST", testServer.URL+"/v1/users/"+testTweetUserID+"/tweets/"+testTweetID, nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("タイムライン情報がデータベース上に登録されること", func(t *testing.T) {
			t.Run("投稿したツイートIDのタイムラインがフォロワー数+自分のタイムラインの数だけレコードが作成されること", func(t *testing.T) {
				rows := handler.Query("SELECT 1 FROM timeline WHERE tweet_id = ?", testTweetID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				var rowCount int
				for rows.Next() {
					rowCount++
				}
				if rowCount != 3 {
					t.Fatalf("rows count unmatched. actual:%d, expected:%d", rowCount, 3)
				}
			})
			t.Run(testTweetUserName+"のタイムラインが生成されていること", func(t *testing.T) {
				rows := handler.Query("SELECT tweet_id FROM timeline WHERE user_id = ?", testTweetUserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				rows.Next()
				var tweetID string
				if err := rows.Scan(&tweetID); err != nil {
					t.Fatalf(err.Error())
				}
				t.Run("tweet_idが"+testTweetID+"のタイムラインが取得できること", func(t *testing.T) {
					if tweetID != testTweetID {
						t.Fatalf("insert faild. actual: %s, expected: %s", tweetID, testTweetID)
					}
				})
			})
			t.Run(testFollower1UserName+"のタイムラインが生成されていること", func(t *testing.T) {
				rows := handler.Query("SELECT tweet_id FROM timeline WHERE user_id = ?", testFollower1UserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				rows.Next()
				var tweetID string
				if err := rows.Scan(&tweetID); err != nil {
					t.Fatalf(err.Error())
				}
				t.Run("tweet_idが"+testTweetID+"のタイムラインが取得できること", func(t *testing.T) {
					if tweetID != testTweetID {
						t.Fatalf("insert faild. actual: %s, expected: %s", tweetID, testTweetID)
					}
				})
			})
			t.Run(testFollower2UserName+"のタイムラインが生成されていること", func(t *testing.T) {
				rows := handler.Query("SELECT tweet_id FROM timeline WHERE user_id = ?", testFollower2UserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				rows.Next()
				var tweetID string
				if err := rows.Scan(&tweetID); err != nil {
					t.Fatalf(err.Error())
				}
				t.Run("tweet_idが"+testTweetID+"のタイムラインが取得できること", func(t *testing.T) {
					if tweetID != testTweetID {
						t.Fatalf("insert faild. actual: %s, expected: %s", tweetID, testTweetID)
					}
				})
			})
		})

	})

	t.Run("POST /v1/users/:userId/follow/:followUserId", func(t *testing.T) {
		testFollowerUserID := uuid.New().String()
		testFollowerUserName := "フォロワーユーザ"
		testFollowUserID := uuid.New().String()
		testFollowUserName := "フォローユーザ"
		testFollowTweet1ID := uuid.New().String()
		testFollowTweet2ID := uuid.New().String()
		testFollowTweet3ID := uuid.New().String()
		// モックデータのINSERT
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollowerUserID, testFollowerUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollowUserID, testFollowUserName, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", testFollowTweet1ID, testFollowUserID, "テストツイート１", "2024-01-01 00:00:01", "2024-01-01 00:00:01")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", testFollowTweet2ID, testFollowUserID, "テストツイート２", "2024-01-01 00:00:02", "2024-01-01 00:00:02")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", testFollowTweet3ID, testFollowUserID, "テストツイート３", "2024-01-01 00:00:03", "2024-01-01 00:00:03")
		handler.Exec("INSERT INTO follow(user_id, follow_user_id, created_at, updated_at) VALUES (?, ?, ?, ?)", testFollowerUserID, testFollowUserID, "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		// タイムライン生成対象外のレコード
		noiseUserID := uuid.New().String()
		noiseTweetID := uuid.New().String()
		handler.Exec("INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)", noiseUserID, "ノイズユーザ", "2024-01-01 00:00:00", "2024-01-01 00:00:00")
		handler.Exec("INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", noiseTweetID, noiseUserID, "ノイズツイート", "2024-01-01 00:00:01", "2024-01-01 00:00:01")

		req, _ := http.NewRequest("POST", testServer.URL+"/v1/users/"+testFollowerUserID+"/follow/"+testFollowUserID, nil)
		httpClient := new(http.Client)
		res, _ := httpClient.Do(req)

		t.Run("レスポンスステータスコードが200であること", func(t *testing.T) {
			if res.StatusCode != 200 {
				t.Fatalf("statusCode Fail. Code: %d, expected: %d", res.StatusCode, 200)
			}
		})

		t.Run("タイムライン情報がデータベース上に登録されること", func(t *testing.T) {
			t.Run("フォローしたユーザのツイートのタイムラインが生成されていること", func(t *testing.T) {
				rows := handler.Query("SELECT t1.tweet_id FROM timeline as t1 INNER JOIN tweet as t2 ON t1.tweet_id = t2.tweet_id WHERE t1.user_id = ? ORDER BY t2.created_at", testFollowerUserID)
				defer func(rows *sql.Rows) {
					if err := rows.Close(); err != nil {
						t.Fatalf(err.Error())
					}
				}(rows)
				var rowCount int
				expectedTweetIDs := []string{
					testFollowTweet1ID, testFollowTweet2ID, testFollowTweet3ID,
				}
				for rows.Next() {
					var tweetID string
					if err := rows.Scan(&tweetID); err != nil {
						t.Fatalf(err.Error())
					}
					t.Run("tweetIDが"+expectedTweetIDs[rowCount]+"と一致すること", func(t *testing.T) {
						if tweetID != expectedTweetIDs[rowCount] {
							t.Fatalf("[%d]tweetID unmatched. actual:%s, expected:%s", rowCount, tweetID, expectedTweetIDs[rowCount])
						}
					})
					rowCount++
				}
				t.Run("フォローしたユーザのツイートの数だけタイムラインのレコードが作成されること", func(t *testing.T) {
					if rowCount != 3 {
						t.Fatalf("rows count unmatched. actual:%d, expected:%d", rowCount, 3)
					}
				})
			})
		})

	})
}
