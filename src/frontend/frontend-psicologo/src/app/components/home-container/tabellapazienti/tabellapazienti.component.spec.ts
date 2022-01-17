import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TabellapazientiComponent } from './tabellapazienti.component';

describe('TabellapazientiComponent', () => {
  let component: TabellapazientiComponent;
  let fixture: ComponentFixture<TabellapazientiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TabellapazientiComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TabellapazientiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
