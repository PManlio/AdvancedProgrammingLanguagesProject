import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AggiungipazientemodalComponent } from './aggiungipazientemodal.component';

describe('AggiungipazientemodalComponent', () => {
  let component: AggiungipazientemodalComponent;
  let fixture: ComponentFixture<AggiungipazientemodalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AggiungipazientemodalComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AggiungipazientemodalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
