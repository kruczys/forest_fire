## Opis programu symulacji spalania lasu

Program symuluje spalanie lasu trafionego piorunem z uwzględnieniem różnych prędkości wiatru i procentowego zagęszczenia drzew. Drzewa są reprezentowane jako elementy dwuwymiarowej tablicy, gdzie każde drzewo posiada określony wiek, który wpływa na jego odporność na ogień. Im młodsze drzewo, tym większa szansa, że zapali się. Symulacja bierze pod uwagę również prędkość wiatru.

## Struktura symulacji

1. **Inicjalizacja lasu** - Losowe umieszczanie drzew z uwzględnieniem zadanej gęstości. Każde drzewo ma przypisany losowy wiek, co wpływa na jego podatność na zapalenie.
2. **Rozpoczęcie pożaru** - Piorun losowo zapala jedno z drzew, inicjując pożar.
3. **Rozprzestrzenianie ognia** - Ogień rozprzestrzenia się od zapalonych drzew do sąsiednich, zależnie od ich wieku i prędkości wiatru.
4. **Obliczenie wyników** - Po zakończeniu symulacji obliczany jest procent spalonych drzew w stosunku do ogólnej liczby drzew.

## Analiza wyników i identyfikacja optymalnego zalesienia

### Wnioski z symulacji:

- **Wpływ prędkości wiatru:** Wzrost prędkości wiatru powoduje szybsze i bardziej intensywne rozprzestrzenianie się pożaru. Przy wysokich wartościach wiatru, pożary są znacznie bardziej destrukcyjne.

### Optymalne zalesienie:

- **Przy niskich prędkościach wiatru (0-250):** Zauważalny skok w poziomie spalenia lasu następuje przy około 60% zalesienia.
- **Przy średnich prędkościach wiatru (250-500):** Próg optymalnego zalesienia to 50%.
- **Przy wysokich prędkościach wiatru (>500):** Optymalne zalesienie spada do 40%.
