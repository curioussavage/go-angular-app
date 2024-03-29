import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { UserGridComponent } from './user-grid/user-grid.component';
import { UserCardComponent } from './user-card/user-card.component';
import { UserFormComponent } from './user-form/user-form.component';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatDialogModule} from '@angular/material/dialog';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { BASE_PATH, DefaultService } from 'projects/api-client';
import { ApiModule } from 'projects/api-client';
import { HttpClientModule } from '@angular/common/http';
import { ConfirmDialogComponent } from './confirm-dialog/confirm-dialog.component';
import { UserEditFormComponent } from './user-edit-form/user-edit-form.component';

@NgModule({
  declarations: [
    AppComponent,
    UserFormComponent,
    UserCardComponent,
    UserGridComponent,
    ConfirmDialogComponent,
    UserEditFormComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserAnimationsModule,
    BrowserModule,
    FormsModule,
    MatButtonModule,
    MatCardModule,
    MatDialogModule,
    MatGridListModule,
    MatIconModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatInputModule,
    ApiModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
    { provide: BASE_PATH, useValue: 'http://localhost:1323/api/v1' },
    DefaultService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
