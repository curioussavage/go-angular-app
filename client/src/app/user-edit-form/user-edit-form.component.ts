import { Component, Input } from '@angular/core';
import { ModelsUser } from 'projects/api-client';
import { Router, ActivatedRoute } from '@angular/router';
import { DefaultService } from 'projects/api-client';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-user-edit-form',
  templateUrl: './user-edit-form.component.html',
  styleUrls: ['./user-edit-form.component.css']
})
export class UserEditFormComponent {
  @Input() user: ModelsUser;
  editForm: FormGroup;

  constructor(private route: ActivatedRoute, private fb: FormBuilder, private router: Router, private defService: DefaultService) {}

  ngOnInit(): void {
    this.editForm = this.fb.group({
      firstName: [],
      lastName: [],
      email: []
    });

    const userId = this.route.snapshot.paramMap.get('id');
    this.defService.usersGet(parseInt(userId || '0')).subscribe(
      (data) => {
        this.user = data[0];

        this.editForm.setValue({
          firstName: this.user.firstName,
          lastName: this.user.lastName,
          email: this.user.email,
        });
      }
    );
  }

  getServerError(field: string): string | undefined {
    const ctl = this.editForm.get(field);
    if (ctl && ctl.errors !== null) {
      return ctl.errors['serverError'];
    }
    return
  }

  updateUser(): void {
    this.defService.userIdPatch(this.editForm.value, this.user.userID!).subscribe(
      (data) => {
        this.router.navigate(['/']);
      },
      (err) => {
        if (Array.isArray(err.error)) {
          for (let e of err.error) {
            const ctl = this.editForm.get(e.field);
            ctl?.setErrors({ serverError: e.message})
          }
        }
      }
    );
  }
}
