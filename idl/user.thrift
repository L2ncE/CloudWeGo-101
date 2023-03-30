namespace go user

struct RegisterRequest {
    1: string username
    2: string password
}

struct ResgisterResponse {}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: i64 uid
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

service UserService {
    LoginResponse Login(1: LoginRequest req)
    ResgisterResponse Register(1: RegisterRequest req)
    GetArticleNumResponse GetArticleNum(1: GetArticleNumRequest req)
    AddArticleNumResponse AddArticleNum(1: AddArticleNumRequest req)
}