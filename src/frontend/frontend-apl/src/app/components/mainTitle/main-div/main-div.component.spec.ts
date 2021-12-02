import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MainDivComponent } from './main-div.component';

describe('MainDivComponent', () => {
  let component: MainDivComponent;
  let fixture: ComponentFixture<MainDivComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MainDivComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MainDivComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
