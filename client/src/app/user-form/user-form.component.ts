import { Component } from '@angular/core';
import {UserSubmission} from '../user';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent {
  user: UserSubmission = { userName: 'asdf', firstName: '', lastName: '', email: '', department: ''};

  onSubmit(): void {
    alert("form submitted")
    this.createUser();
  }

  createUser(): void {
    console.log("creating the user")
  }

}
