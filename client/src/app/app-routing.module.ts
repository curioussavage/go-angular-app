import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserFormComponent } from './user-form/user-form.component';
import { UserEditFormComponent } from './user-edit-form/user-edit-form.component';
import { UserGridComponent } from './user-grid/user-grid.component';

const routes: Routes = [
  { path: 'users', component: UserGridComponent },
  { path: 'new-user', component: UserFormComponent },
  { path: 'edit-user/:id', component: UserEditFormComponent },
  { path: '', redirectTo: '/users', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
