
export enum Status {
  INACTIVE = 'I',
  ACTIVE = 'A',
  TERMINATED = 'T',
}

export interface User {
  userId: number
  userName: string
  firstName: string
  lastName: string
  email: string
  userStatus: Status
  department: string
}

export interface UserSubmission {
  userName: string
  firstName: string
  lastName: string
  email: string
  department: string
}
