export interface User {
  id: string
  email: string
  first_name: string
  last_name: string
  role: string
  email_verified: boolean
}

export interface AuthResponse {
  user: User
  access_token: string
  refresh_token: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  first_name: string
  last_name: string
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  message?: string
  error?: {
    code: string
    message: string
    details?: Array<{
      field: string
      message: string
    }>
  }
}
