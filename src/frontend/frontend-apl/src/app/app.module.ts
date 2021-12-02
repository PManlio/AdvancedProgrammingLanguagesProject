import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { MainDivComponent } from './components/mainTitle/main-div/main-div.component';
import { MainLoginComponent } from './components/mainTitle/main-div/main-login/main-login.component';
import { MainRegisterComponent } from './components/mainTitle/main-div/main-register/main-register.component';

@NgModule({
  declarations: [
    AppComponent,
    MainDivComponent,
    MainLoginComponent,
    MainRegisterComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
