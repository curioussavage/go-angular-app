import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from '../app.module';

import { MatDialogModule, MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';

import { ConfirmDialogComponent } from './confirm-dialog.component';

import { MatDialogRef } from '@angular/material/dialog';


describe('ConfirmDialogComponent', () => {
  let component: ConfirmDialogComponent;
  let fixture: ComponentFixture<ConfirmDialogComponent>;

  let closeSpy: jasmine.Spy
  beforeEach(async () => {
    closeSpy = jasmine.createSpy('close');
    await TestBed.configureTestingModule({
      imports: [AppModule, MatDialogModule],
      declarations: [ ConfirmDialogComponent ],
      providers: [
        {provide: MAT_DIALOG_DATA, useValue: { message: "test" }},
        {provide: MatDialogRef, useValue: { close: closeSpy }}
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ConfirmDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', async () => {
    expect(component).toBeTruthy();
  });

  it('should display the message', async () => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.textContent)
      .toContain("test");
  });

  it('should close dialog on confirm clicked', () => {
    component.onConfirm();
    fixture.detectChanges();
    expect(closeSpy.calls.count()).toBe(1, 'dialog closed');
  });

  it('should close dialog on Cancel clicked', () => {
    component.onCancel();
    fixture.detectChanges();
    expect(closeSpy.calls.count()).toBe(1, 'dialog closed');
  });
});
