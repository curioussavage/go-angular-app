import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from '../app.module';

import {MatDialogHarness} from '@angular/material/dialog/testing';
import { MatDialogModule, MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';

import { ConfirmDialogComponent } from './confirm-dialog.component';

import { MatDialogRef } from '@angular/material/dialog';


describe('ConfirmDialogComponent', () => {
  let component: ConfirmDialogComponent;
  let fixture: ComponentFixture<ConfirmDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AppModule, MatDialogModule],
      declarations: [ ConfirmDialogComponent ],
      providers: [
        {provide: MAT_DIALOG_DATA, useValue: { message: "test" }},
        {provide: MatDialogRef, useValue: { close: (res: any) => res }}
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ConfirmDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', async () => {
    MatDialogHarness
    expect(component).toBeTruthy();
  });
});
