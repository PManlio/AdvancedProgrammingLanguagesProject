import { TestBed } from '@angular/core/testing';

import { PazienteServiceService } from './paziente-service.service';

describe('PazienteServiceService', () => {
  let service: PazienteServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PazienteServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
