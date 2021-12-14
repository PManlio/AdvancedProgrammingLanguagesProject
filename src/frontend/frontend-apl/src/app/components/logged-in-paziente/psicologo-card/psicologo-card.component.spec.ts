import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PsicologoCardComponent } from './psicologo-card.component';

describe('PsicologoCardComponent', () => {
  let component: PsicologoCardComponent;
  let fixture: ComponentFixture<PsicologoCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PsicologoCardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PsicologoCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
