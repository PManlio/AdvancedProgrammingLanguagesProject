import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './modules/app-routing/app-routing.module';

import { AppComponent } from './app.component';
import { MainDivComponent } from './components/mainTitle/main-div/main-div.component';
import { MainLoginComponent } from './components/mainTitle/main-div/main-login/main-login.component';
import { MainRegisterComponent } from './components/mainTitle/main-div/main-register/main-register.component';
import { LogoutbtnComponent } from './components/logoutbtn/logoutbtn.component';
import { RegistrationModalComponent } from './components/mainTitle/main-div/registration-modal/registration-modal.component';
import { LoggedInPazienteComponent } from './components/logged-in-paziente/logged-in-paziente.component';
import { PsicologoCardComponent } from './components/logged-in-paziente/psicologo-card/psicologo-card.component';
import { ScriviDiarioComponent } from './components/logged-in-paziente/scrivi-diario/scrivi-diario.component';
import { CercapsicologoComponent } from './components/logged-in-paziente/cercapsicologo/cercapsicologo.component';
import { PsicologotrovatoComponent } from './components/logged-in-paziente/cercapsicologo/psicologotrovato/psicologotrovato.component';

@NgModule({
  declarations: [
    AppComponent,
    MainDivComponent,
    MainLoginComponent,
    MainRegisterComponent,
    LogoutbtnComponent,
    RegistrationModalComponent,
    LoggedInPazienteComponent,
    PsicologoCardComponent,
    ScriviDiarioComponent,
    CercapsicologoComponent,
    PsicologotrovatoComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
