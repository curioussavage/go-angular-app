import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from '../app.module';

import { UserCardComponent } from './user-card.component';

describe('UserCardComponent', () => {
  let component: UserCardComponent;
  let fixture: ComponentFixture<UserCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AppModule],
      declarations: [ UserCardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UserCardComponent);
    component = fixture.componentInstance;
    component.user = { userName: 'foo', firstName: "john", lastName: "doe", email: "foo@bar.com", department: "" }
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display the username', () => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('#username').textContent)
      .toContain(component.user.userName);
  });
});
