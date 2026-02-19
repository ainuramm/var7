# Lab 4 — Variant 7: FuelTrip

## Цель

1. Создать собственный пакет `fueltrip` с экспортируемыми функциями.
2. Подключить пакет из GitHub и использовать его в `main.go`.
3. Использовать внешние пакеты `uuid` и `color`.
4. Реализовать обработку ошибок, работу с указателем и форматированный вывод.
5. Проверить документацию через `go doc`.

---

## Структура проекта
var7/
├─ go.mod
├─ README.md
├─ cmd/
│  └─ app/
│     └─ main.go
└─ pkg/
   └─ fueltrip/
      └─ fueltrip.go
## Функции пакета fueltrip

1. `FuelNeeded(distanceKm, litersPer100 float64) (float64, error)`  
   - Вычисляет количество топлива для поездки.

2. `TripCost(fuelLiters, pricePerLiter float64) (float64, error)`  
   - Считает стоимость поездки.

3. `AddTrafficReserve(fuel *float64, percent float64) error`  
   - Добавляет резерв топлива в процентах.

4. `FormatTripReport(route string, fuel, cost float64) (string, error)`  
   - Формирует строку отчёта о поездке.
