import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { PsicologoCardComponent } from 'src/app/components/logged-in-paziente/psicologo-card/psicologo-card.component';
import { ScriviDiarioComponent } from 'src/app/components/logged-in-paziente/scrivi-diario/scrivi-diario.component';
import { CercapsicologoComponent } from 'src/app/components/logged-in-paziente/cercapsicologo/cercapsicologo.component';


const routes: Routes = [
  { path: '', component: PsicologoCardComponent },
  { path: 'scrividiario', component: ScriviDiarioComponent },
  { path: 'cercapsicologo', component: CercapsicologoComponent }
]

@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
