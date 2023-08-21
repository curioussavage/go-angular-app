import { Injectable } from '@angular/core';
import { User, Status } from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor() { }

  getUsers(): User[] {
    return [
      {
        userId: 1,
        userName: 'foo',
        firstName: 'bob',
        lastName: 'jones',
        email: 'foo@b.com',
        userStatus: Status.ACTIVE,
        department: 'fun'
      },
      {
        userId: 1,
        userName: 'bar_123',
        firstName: 'bob',
        lastName: 'jones',
        email: 'foo@b.com',
        userStatus: Status.ACTIVE,
        department: 'fun'
      }
    ]
  }
}
