import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CercapsicologoComponent } from './cercapsicologo.component';

describe('CercapsicologoComponent', () => {
  let component: CercapsicologoComponent;
  let fixture: ComponentFixture<CercapsicologoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CercapsicologoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CercapsicologoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
