import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InfopsicologoComponent } from './infopsicologo.component';

describe('InfopsicologoComponent', () => {
  let component: InfopsicologoComponent;
  let fixture: ComponentFixture<InfopsicologoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InfopsicologoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InfopsicologoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
