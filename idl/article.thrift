namespace go article

struct BaseResponse {
    1: i64 status_code,
    2: string status_msg,
}

struct PostArticleRequest {
    1: string title
    2: string content
}

struct PostArticleResponse {
    1: BaseResponse base
}

service ArticleService {
    PostArticleResponse PostArticle(1: PostArticleResponse req)
}