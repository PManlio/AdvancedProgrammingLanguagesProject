import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogoutbtnComponent } from './logoutbtn.component';

describe('LogoutbtnComponent', () => {
  let component: LogoutbtnComponent;
  let fixture: ComponentFixture<LogoutbtnComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LogoutbtnComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LogoutbtnComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
