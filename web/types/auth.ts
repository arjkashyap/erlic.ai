export interface NullableString {
  String: string
  Valid: boolean
}

export interface User {
  Id: number
  FirstName: string
  LastName: string
  FullName: string
  Username: string
  Email: string
  Organization: NullableString
  Provider: string
  ProviderUserID: string
  AvatarURL: NullableString
}

export interface AuthResponse {
  authenticated: boolean
  user: User | null
} 