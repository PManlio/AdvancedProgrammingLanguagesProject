import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ScriviDiarioComponent } from './scrivi-diario.component';

describe('ScriviDiarioComponent', () => {
  let component: ScriviDiarioComponent;
  let fixture: ComponentFixture<ScriviDiarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ScriviDiarioComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ScriviDiarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
