import { Component } from '@angular/core';
import { Router } from '@angular/router';
import {UserSubmission} from '../user';
import { DefaultService, ModelsUserCreationForm } from 'projects/api-client';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent {
  user: UserSubmission = { userName: '', firstName: '', lastName: '', email: '', department: ''};
  form: FormGroup;

  constructor(private fb: FormBuilder, private router: Router, private defService: DefaultService) {}

  ngOnInit() {
    this.form = this.fb.group({
      userName: [],
      firstName: [],
      lastName: [],
      email: [],
      department: [],
    });
  }

  onSubmit(): void {
    this.createUser();
  }

  getServerError(field: string): string | undefined {
    const ctl = this.form.get(field);
    if (ctl && ctl.errors !== null) {
      return ctl.errors['serverError'];
    }
    return
  }

  createUser(): void {
    this.defService.userPost(this.form.value).subscribe(
      (data) => {
        console.log("got data");
        this.router.navigate(['/']);
      },
      (err) => {
        if (Array.isArray(err.error)) {
          for (let e of err.error) {
            const ctl = this.form.get(e.field);
            ctl?.setErrors({ serverError: e.message})
          }
        }
      }
    );
  }

}
