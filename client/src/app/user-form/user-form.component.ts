import { Component } from '@angular/core';
import { Router } from '@angular/router';
import {UserSubmission} from '../user';
import { DefaultService, ModelsUserCreationForm } from 'projects/api-client';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent {
  user: UserSubmission = { userName: '', firstName: '', lastName: '', email: '', department: ''};

  constructor(private router: Router, private defService: DefaultService) {}

  onSubmit(): void {
    this.createUser();
  }

  createUser(): void {
    this.defService.userPost(this.user).subscribe(
      (data) => {
        console.log("got data");
        this.router.navigate(['/']);
      }
    )
    console.log("creating the user")
  }

}
