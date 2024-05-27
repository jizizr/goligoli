namespace go user

struct RegisterRequest {
    1: string username,
    2: string password,
}

struct RegisterResponse {
    2: string token,
}

struct LoginRequest {
    1: string username,
    2: string password,
}

struct LoginResponse {
    1: string token,
}

service UserService {
    LoginResponse Login(1: LoginRequest req),
    RegisterResponse Register(1: RegisterRequest req),
}
