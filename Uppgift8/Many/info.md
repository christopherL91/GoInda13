#Många sändare och många mottagare

Programmet many2many.go innehåller 4 producenter som tillsammans skickar 32 strängar över en kanal; i andra änden av kanalen finns 2 konsumenter som tar emot de 32 strängarna. Förklara vad som händer och varför det händer om man gör följande ändringar i programmet. Prova att först tänka ut vad som händer och testa sedan din hypotes genom att ändra och köra programmet.

* Vad händer om man byter plats på satserna wgp.Wait() och close(ch) i slutet av main-funktionen?

* Vad händer om man flyttar close(ch) från main-funktionen och i stället stänger kanalen i slutet av funktionen Produce?

* Vad händer om man tar bort satsen close(ch) helt och hållet?

* Vad händer om man ökar antalet konsumenter från 2 till 4?

* Kan man vara säker på att alla strängar blir utskrivna innan programmet stannar?
