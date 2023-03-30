namespace go api

// User

struct RegisterRequest {
    1: string username
    2: string password
}

struct ResgisterResponse {
    1: string token
}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string token
}

struct GetArticleNumRequest {
    1: i64 user_id
}

struct GetArticleNumResponse {
    1: i64 num
}

struct AddArticleNumRequest {
    1: i64 user_id
}

struct AddArticleNumResponse {}

// Article

struct PostArticleRequest {
    1: string title
    2: string content
}

struct PostArticleResponse {}

service ApiService {
    ResgisterResponse Register(1: RegisterRequest req)(api.post="/user/register/");
    LoginResponse Login(1: LoginRequest req)(api.post="/user/login/");
    GetArticleNumResponse GetArticleNum(1: GetArticleNumRequest req)(api.get="/user/article");
    AddArticleNumResponse AddArticleNum(1: AddArticleNumRequest req)(api.post="/user/article");

    PostArticleResponse PostArticle(1: PostArticleRequest req)(api.post="article");
}