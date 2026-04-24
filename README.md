# Projelerini sergile, yeteneklerini göster ve keşfet! (Örnek)

<div align="center">

[![GitHub](https://img.shields.io/badge/GitHub-Bilgisayar-Kavramlari-Toplulugu-181717?style=flat-square&logo=github)](https://github.com/Bilgisayar-Kavramlari-Toplulugu/project-devboard)
# 🖥️ DevBoard

Projelerini sergile, yeteneklerini göster ve keşfet!

[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?style=flat-square&logo=postgresql)](https://www.postgresql.org)
[![GoLang](https://img.shields.io/badge/GO-1.25-007d9c?style=flat-square&logo=go)](https://go.dev/)
[![Vue 3](https://img.shields.io/badge/Vue-3-42b883?style=flat-square&logo=vue)](https://vuejs.org/)
<br>
[![GitHub](https://img.shields.io/badge/GitHub-Bilgisayar--Kavramlari--Toplulugu-181717?style=flat-square&logo=github)](https://github.com/Bilgisayar-Kavramlari-Toplulugu/project-devboard)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](LICENSE)

**Part of [Projelerini sergile, yeteneklerini göster ve keşfet!](docs/Project-Definition.md)**
**Part of [Bilgisayar Kavramları Topluluğu](https://github.com/Bilgisayar-Kavramlari-Toplulugu)**

</div>


<br>

> **ÖNEMLİ:** Bu repository **Projelerini sergile, yeteneklerini göster ve keşfet!** projesinin bir parçasıdır. Proje hakkında detaylı bilgi için [`docs/Project-Definition.md`](docs/Project-Definition.md) dosyasına bakın.
> **ÖNEMLİ:** Bu repository **DevBoard** projesinin bir parçasıdır. Proje hakkında detaylı bilgi için [`docs/Project-Definition.md`](docs/Project-Definition.md) dosyasına bakın.

## 📖 Hakkında

**DevBoard**, yazılım geliştiricilerin projelerini ve yetkinliklerini sergilediği, çalışan veya takım arkadaşı arayanların ise nitelikli adayları filtreleyerek keşfettiği dinamik bir ekosistemdir. Bilgisayar Kavramları (BK) bünyesindeki projeleri temel alan referans ve onaylama (endorsement) sistemi sayesinde, geliştiricilerin gerçek dünya deneyimleri doğrulanabilir ve şeffaf bir şekilde sunulur.
<!-- Bu repository'nin ne yaptığını buraya yazın -->

## 🚀 Kurulum

### Gereksinimler

- Gerekli araçları buraya listeleyin
- [Git](https://git-scm.com/)

### Başlangıç

> **IMPORTANT:** This repository is part of **Projelerini sergile, yeteneklerini göster ve keşfet!** project. See [`docs/Project-Definition.md`](docs/Project-Definition.md) for details.

## 📖 About

**DevBoard** is a dynamic ecosystem where software developers showcase their projects and competencies, and recruiters discover qualified candidates through advanced filtering. Leveraging a reference and endorsement system based on "Bilgisayar Kavramları" (BK) projects, it ensures that developers' real-world experiences are presented in a verifiable and transparent manner.
<!-- Describe what this repository does -->

## 🚀 Installation

### Requirements

- List required tools here
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org)
- [GO](https://go.dev/doc/install)
- [Vue](https://vuejs.org/)

### Getting Started


## 🔐 .env Yönetimi

### Ortam Bazlı Yaklaşım

DevBoard'da ortam değişkenleri ortama göre farklı yönetilir:

| Ortam | Yönetim | Dosya | Açıklama |
|-------|---------|-------|----------|
| **Development** | Lokal `.env.age` → Decrypt | `.env` | Geliştirici tarafından decrypt edilir |
| **Staging** | Google Secrets Manager | Yok | GCP Console'dan yönetilir |
| **Production** | Google Secrets Manager | Yok | GCP Console'dan yönetilir |

⚠️ **Önemli**: Staging ve production ortamları için lokal `.env` dosyası oluşturmanıza gerek yoktur. Bu ortamların değişkenlerini değiştirmek için DevOps ekibine başvurun.

### Lokal Geliştirme Kurulumu

#### Gereksinimler

- **age** şifreleme aracı
- GitHub hesabında SSH Authentication key

#### 1. age Aracını Kurun

```bash
# macOS
brew install age

# Linux (Debian/Ubuntu)
apt install age

# Windows (PowerShell)
winget install FiloSottile.age

# Doğrula
age --version
```

#### 2. GitHub SSH Key'i Ekleyin

GitHub hesabınızda SSH Authentication key yoksa ekleyin:

```bash
# Mevcut key'leri kontrol et
ls ~/.ssh/

# Key yoksa oluştur
ssh-keygen -t ed25519 -C "your-email@example.com"

# Public key'i göster
cat ~/.ssh/id_ed25519.pub
```

Çıktıyı kopyalayıp GitHub Settings → SSH and GPG keys → New SSH key olarak ekleyin.

#### 3. .env.age Dosyasını Decrypt Edin

Proje kök dizininde çalıştırın:

```bash
./encrypt-env-file.sh decrypt
```

Bu komut `.env.age` dosyasını decrypt ederek `.env` dosyası oluşturur.

Doğrulama:
```bash
cat .env | grep ENVIRONMENT  # Development ortamını göstermelidir
```

### encrypt-env-file.sh Script Referansı

Bu script, `.env` dosyasını GitHub SSH key'leri kullanarak şifreler/çözer.

**Şifreleme (Encrypt)** — `.env` → `.env.age`:
```bash
./encrypt-env-file.sh
```

**Şifre Çözme (Decrypt)** — `.env.age` → `.env`:
```bash
./encrypt-env-file.sh decrypt
```

### Güvenlik Notları

- ✅ `.env` dosyası Git'e commit edilmez
- ✅ Yalnızca `.env.age` Git'e eklenebilir
- ✅ Şifreleme GitHub username'lere bağlıdır
- ✅ Private key'iniz gizli kalmalıdır

### Alıcı (Recipient) Yönetimi

Şifreleme yetkisi olan kullanıcıları değiştirmek için `encrypt-env-file.sh` scriptinin başındaki `RECIPIENTS` listesini düzenleyin. Listeye eklenen her GitHub username:
- GitHub'daki SSH public key'i ile şifrelemeye dahil edilir
- Kendi private key'i ile dosyayı decrypt edebilir

