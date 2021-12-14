import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoggedInPazienteComponent } from './logged-in-paziente.component';

describe('LoggedInPazienteComponent', () => {
  let component: LoggedInPazienteComponent;
  let fixture: ComponentFixture<LoggedInPazienteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoggedInPazienteComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoggedInPazienteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
