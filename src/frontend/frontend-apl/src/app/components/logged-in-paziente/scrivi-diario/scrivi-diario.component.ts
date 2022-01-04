import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { DiaryService } from 'src/app/services/diary.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-scrivi-diario',
  templateUrl: './scrivi-diario.component.html',
  styleUrls: ['./scrivi-diario.component.css']
})
export class ScriviDiarioComponent implements OnInit, OnDestroy {

  private subscriptionToRouteParams: any;
  public emailPsicologo: number;
  private email: string;

  constructor(private router: ActivatedRoute, private userInfo: UserInfoService, private diaryService: DiaryService) {
    this.userInfo.email.subscribe(email => this.email = email);
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
    this.diaryService.postDiary(this.email, text)
  }
}
