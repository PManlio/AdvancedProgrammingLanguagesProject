import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-scrivi-diario',
  templateUrl: './scrivi-diario.component.html',
  styleUrls: ['./scrivi-diario.component.css']
})
export class ScriviDiarioComponent implements OnInit, OnDestroy {

  private subscriptionToRouteParams: any;
  public emailPsicologo: number;

  constructor(private router: ActivatedRoute) { }

  ngOnInit(): void {
    this.subscriptionToRouteParams = this.router.params.subscribe( params => {
      this.emailPsicologo = params['emailPsicologo']
    })
  }

  ngOnDestroy(): void {
    this.subscriptionToRouteParams.unsubscribe();
  }
}
