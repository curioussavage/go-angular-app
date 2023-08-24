import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppModule } from '../app.module';

import { UserGridComponent } from './user-grid.component';

describe('UserGridComponent', () => {
  let component: UserGridComponent;
  let fixture: ComponentFixture<UserGridComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AppModule],
      declarations: [ UserGridComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UserGridComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
