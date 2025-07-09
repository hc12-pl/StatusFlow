# StatusFlow

**StatusFlow** to rozwijany projekt open-source w języku Go, który ma na celu uproszczenie lokalnego monitoringu podstawowych usług sieciowych (HTTP, TCP, DNS, ICMP) bez potrzeby korzystania z rozbudowanych systemów typu Zabbix czy Prometheus.

Projekt jest w fazie rozwoju – zapraszam do śledzenia postępów i współtworzenia!

## ✨ Cechy (planowane / częściowo gotowe)
- Lekki, lokalny monitoring zasobów
- GUI oparte na bibliotece [Fyne](https://fyne.io/)
- Brak konieczności użycia backendu, bazy danych czy kontenerów
- Możliwość konfiguracji hostów do monitorowania

## 🔍 W planie testy:
- [x] HTTP
- [x] TCP
- [ ] DNS
- [ ] ICMP (ping)
- [ ] Harmonogram testów
- [ ] Eksport logów
- [ ] System powiadomień

## 🚀 Jak uruchomić
Projekt wymaga środowiska Go:

```bash
git clone https://github.com/szymonwys/StatusFlow
cd StatusFlow
go run cmd/main.go
