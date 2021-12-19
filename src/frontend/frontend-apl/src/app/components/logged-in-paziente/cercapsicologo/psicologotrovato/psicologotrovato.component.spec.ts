import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PsicologotrovatoComponent } from './psicologotrovato.component';

describe('PsicologotrovatoComponent', () => {
  let component: PsicologotrovatoComponent;
  let fixture: ComponentFixture<PsicologotrovatoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PsicologotrovatoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PsicologotrovatoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
