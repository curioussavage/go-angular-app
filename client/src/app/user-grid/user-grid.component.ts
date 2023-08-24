import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { DefaultService, ModelsUser } from 'projects/api-client';

@Component({
  selector: 'app-user-grid',
  templateUrl: './user-grid.component.html',
  styleUrls: ['./user-grid.component.css'],
})
export class UserGridComponent {
  users?: Observable<ModelsUser[]>;

  constructor(private defService: DefaultService) {}

  getUsers(): void {
    this.users = this.defService.usersGet()
  }

  ngOnInit(): void {
    this.getUsers();
  }

  onDelete(userId: number) {
    console.log("onDelete ", userId)
    this.defService
      .userIdDelete(userId)
      .subscribe(
        (resp) => {
          this.getUsers();
        },
        (err) => {
          // TODO
        }
      );
  }
}
