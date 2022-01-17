import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeContainerComponent } from './components/home-container/home-container.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { LogoutBtnComponent } from './components/logout-btn/logout-btn.component';
import { TabellapazientiComponent } from './components/home-container/tabellapazienti/tabellapazienti.component';
import { InfopsicologoComponent } from './components/home-container/infopsicologo/infopsicologo.component';
import { DiarioComponent } from './components/home-container/diario/diario.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeContainerComponent,
    LoginComponent,
    RegisterComponent,
    LogoutBtnComponent,
    TabellapazientiComponent,
    InfopsicologoComponent,
    DiarioComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
