namespace go article

struct PostArticleRequest {
    1: string title
    2: string content
    3: i64 uid
}

struct PostArticleResponse {}

service ArticleService {
    PostArticleResponse PostArticle(1: PostArticleRequest req)
}