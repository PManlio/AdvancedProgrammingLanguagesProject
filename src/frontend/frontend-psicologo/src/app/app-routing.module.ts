import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DiarioComponent } from './components/home-container/diario/diario.component';
import { TabellapazientiComponent } from './components/home-container/tabellapazienti/tabellapazienti.component';

const routes: Routes = [
  { path:'', component: TabellapazientiComponent },
  { path:'diario', component: DiarioComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
