import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-scrivi-diario',
  templateUrl: './scrivi-diario.component.html',
  styleUrls: ['./scrivi-diario.component.css']
})
export class ScriviDiarioComponent implements OnInit, OnDestroy {

  private subscriptionToRouteParams: any;
  public emailPsicologo: number;
  private codFisc: string;

  constructor(private router: ActivatedRoute, private userInfo: UserInfoService) {
    this.userInfo.codFisc.subscribe(codFisc => this.codFisc = codFisc);
  }

  ngOnInit(): void {
    this.subscriptionToRouteParams = this.router.params.subscribe( params => {
      this.emailPsicologo = params['emailPsicologo']
    })
  }

  ngOnDestroy(): void {
    this.subscriptionToRouteParams.unsubscribe();
  }

  public sendText(text: any) {
    // inserire qui la query
    console.log(this.emailPsicologo, this.codFisc, text)
  }
}
