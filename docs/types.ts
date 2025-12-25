// types.ts

export interface User {
  id: string;
  username: string;
  email: string;
  password: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface AuthToken {
  token: string;
  expiresAt: Date;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface ErrorResponse {
  error: string;
  message: string;
  statusCode: number;
}

export enum UserRole {
  ADMIN = 'admin',
  USER = 'user',
}

export interface JWTToken {
  sub: string;
  role: UserRole;
  exp: number;
  iat: number;
}