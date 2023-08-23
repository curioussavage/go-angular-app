import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../user';
import { UserService } from '../user.service';
import { DefaultService, ModelsUser } from 'projects/api-client';

@Component({
  selector: 'app-user-grid',
  templateUrl: './user-grid.component.html',
  styleUrls: ['./user-grid.component.css'],
})
export class UserGridComponent {
  users: Observable<ModelsUser[]>;

  constructor(private userService: UserService, private defService: DefaultService) {}

  getUsers(): void {
    this.users = this.defService.usersGet()
    // this.users = this.userService.getUsers();
  }

  ngOnInit(): void {
    this.getUsers();
  }
}
