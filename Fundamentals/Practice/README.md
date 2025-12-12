# 🧠 GO AMALIY TOPSHIRIQ

## Mavzu: **Concurrent Bank Transaction System**

### 🎯 Maqsad

Go tilida **real hayotga yaqin bank tizimi** yaratish. Dastur bir vaqtning o‘zida bir nechta tranzaksiyalarni bajarishi, xatolarni to‘g‘ri qaytarishi va turli xil bildirishnomalarni (notification) qo‘llab-quvvatlashi kerak.

---

## 📌 Umumiy talablar

Dastur quyidagi **Go mavzularini majburiy** ishlatishi shart:

* `struct`
* `pointer`
* `interface`
* `method receiver`
* `goroutine`
* `channel`
* `mutex`
* `error handling`
* `data type handling (type switch yoki type assertion)`

---

## 1️⃣ Structlar

### 👤 User

* Foydalanuvchini ifodalaydi
* Kamida bitta maydonga ega bo‘lishi kerak (masalan: ism)

### 💳 Account

* Bank hisobini ifodalaydi
* Quyidagilar bo‘lishi shart:

  * Hisob egasi (User ga pointer)
  * Balans (integer)
  * Mutex (concurrency uchun)

### 💼 Transaction

* Bitta tranzaksiyani ifodalaydi
* Quyidagilarni o‘z ichiga olishi kerak:

  * Tranzaksiya turi (deposit yoki withdraw)
  * Tranzaksiya summasi

---

## 2️⃣ Pointer bilan ishlash

* Account balansini o‘zgartiradigan **barcha methodlar pointer receiver (`*Account`) orqali yozilishi shart**
* Balans real vaqtda o‘zgarishi kerak

---

## 3️⃣ Error handling

Dastur quyidagi xatolarni aniqlab, error qaytarishi kerak:

* Noto‘g‘ri summa (0 yoki manfiy son)
* Yetarli mablag‘ yo‘qligi
* Noma’lum tranzaksiya turi

Custom errorlardan foydalanish **majburiy**.

---

## 4️⃣ Interface (Notification tizimi)

### 🔔 Notifier interface

* Kamida bitta methodga ega bo‘lishi kerak (masalan: xabar yuborish)

### Interface’ni amalga oshiruvchi turlar:

* Kamida **2 xil notifier** bo‘lishi shart
  (masalan: Email, SMS, Telegram va hokazo)

---

## 5️⃣ Data type handling

* Notification yuborishda **type switch yoki type assertion** ishlatilishi kerak
* Qaysi notifier ishlatilayotganini runtime’da aniqlash kerak

---

## 6️⃣ Concurrency (majburiy qism)

* Tranzaksiyalar **channel orqali** uzatilishi kerak
* Kamida **1 ta goroutine** tranzaksiyalarni qayta ishlashi kerak
* Bir vaqtning o‘zida keladigan tranzaksiyalar **race condition** keltirib chiqarmasligi kerak (mutex shart)

---

## 7️⃣ Dastur oqimi (mantiq)

1. Foydalanuvchi va bank hisobi yaratiladi
2. Bir nechta tranzaksiya channel orqali yuboriladi
3. Goroutine tranzaksiyalarni ketma-ket qayta ishlaydi
4. Har bir muvaffaqiyatli yoki xato holatda notification yuboriladi
5. Dastur oxirida yakuniy balans chiqariladi

---

## 8️⃣ Natija (kutilayotgan xulq)

* Dastur **panic qilmasligi** kerak
* Barcha xatolar **error orqali** boshqarilishi kerak
* Bir vaqtning o‘zida bir nechta tranzaksiya ishlaganda balans buzilmasligi kerak
* Interface yordamida turli notifierlar bir xil joyda ishlashi kerak

---

## 🔥 Qo‘shimcha (ixtiyoriy, lekin kuchli PLUS)

Agar xohlasang, quyidagilarni ham qo‘shishing mumkin:

* Tranzaksiya tarixini saqlash
* Bir nechta worker goroutine
* `sync.WaitGroup` ishlatish
* JSON’dan tranzaksiyalarni o‘qish
* Timeout (`context.Context`) qo‘llash

---

## ✅ Baholash mezoni (o‘zingni tekshirish uchun)

* Pointer nega ishlatilganini tushuntira olasanmi?
* Mutex olib tashlansa nima bo‘lishini bilasanmi?
* Interface’ning foydasini real misolda ko‘rsata oldingmi?
* Error va oddiy print farqini tushunyapsanmi?


Qachon tayyor bo‘lsa — kodni tashla 👌
