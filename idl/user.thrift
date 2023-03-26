namespace go user

struct BaseResponse {
    1: i64 status_code,
    2: string status_msg,
}

struct RegisterRequest {
    1: string username
    2: string password
}

struct ResgisterResponse {
    1: string token
    2: BaseResponse base
}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: string token
    2: BaseResponse base
}

struct GetArticleNumRequest {
    1: i64 user_id
}

struct GetArticleNumResponse {
    1: i64 num
    2: BaseResponse base
}

struct AddArticleNumRequest {
    1: i64 user_id
}

struct AddArticleNumResponse {
    1: BaseResponse base
}

service UserService {
    LoginResponse Login(1: LoginRequest req)
    ResgisterResponse Register(1: RegisterRequest req)
    GetArticleNumResponse GetArticleNum(1: GetArticleNumRequest req)
    AddArticleNumResponse AddArticleNum(1: AddArticleNumRequest req)
}