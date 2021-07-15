export interface User {
  email: string;
}

export interface UserSubmit {
  email: string;
  password: string;
}

export interface UserResponse {
  user: User;
}
