import { Component } from '@angular/core';
import { User } from '../user';
import { UserService } from '../user.service';

@Component({
  selector: 'app-user-grid',
  templateUrl: './user-grid.component.html',
  styleUrls: ['./user-grid.component.css'],
})
export class UserGridComponent {
  users: User[] = [];

  constructor(private userService: UserService) {}

  getUsers(): void {
    this.users = this.userService.getUsers();
  }

  ngOnInit(): void {
    this.getUsers();
  }
}
