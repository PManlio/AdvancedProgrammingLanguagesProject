import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RimuovipazienteBtnComponent } from './rimuovipaziente-btn.component';

describe('RimuovipazienteBtnComponent', () => {
  let component: RimuovipazienteBtnComponent;
  let fixture: ComponentFixture<RimuovipazienteBtnComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RimuovipazienteBtnComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RimuovipazienteBtnComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
