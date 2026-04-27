## 🔐 .env Yönetimi

### Ortam Bazlı Yaklaşım

DevBoard'da ortam değişkenleri ortama göre farklı yönetilir:

| Ortam | Yönetim | Dosya | Açıklama |
|-------|---------|-------|----------|
| **Development** | Lokal `.env.age` → Decrypt | `.env` | Geliştirici tarafından decrypt edilir |

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


# Bash Script to manage .env file (encrypt-decrypt-env-file.sh)
Copy this script into your project.

```bash
#!/bin/bash

# ── Usage ─────────────────────────────────────────────────────────────────────
#
#   Encrypt .env → .env.age   (run from the directory containing .env):
#     ./encrypt-env-file.sh
#
#   Decrypt .env.age → .env   (run from the directory containing .env.age):
#     ./encrypt-env-file.sh decrypt
#
#   Requirements:
#     - 'age' must be installed  (macOS: brew install age | Windows: winget install FiloSottile.age)
#     - Your SSH public key must be registered as a recipient to decrypt
#
#   To add or remove recipients:
#     Edit the RECIPIENTS list near the top of this script
#
# ─────────────────────────────────────────────────────────────────────────────

# ── Recipients ───────────────────────────────────────────────────────────────
# Add or remove GitHub usernames here to manage encryption recipients
RECIPIENTS=(
    ysfcc
    flovearth
    # add new participants here

)
# ─────────────────────────────────────────────────────────────────────────────

# ── Check if age is installed ─────────────────────────────────────────────────
if ! command -v age &> /dev/null; then
    echo "✗ 'age' is not installed. Install it first:"
    echo ""
    echo "  macOS:   brew install age"
    echo ""
    echo "  Windows: winget install FiloSottile.age"
    echo "           (or download from https://github.com/FiloSottile/age/releases)"
    echo ""
    exit 1
fi
# ─────────────────────────────────────────────────────────────────────────────

# ── Decrypt mode ─────────────────────────────────────────────────────────────
if [ "${1}" = "decrypt" ]; then
    if [ ! -f .env.age ]; then
        echo "✗ .env.age not found in current directory"
        exit 1
    fi

    # Find default SSH private key
    SSH_KEY="${HOME}/.ssh/id_ed25519"
    if [ ! -f "${SSH_KEY}" ]; then
        SSH_KEY="${HOME}/.ssh/id_rsa"
    fi
    if [ ! -f "${SSH_KEY}" ]; then
        echo "✗ No SSH private key found at ~/.ssh/id_ed25519 or ~/.ssh/id_rsa"
        echo "  Specify a key manually: age -d -i /path/to/key .env.age > .env"
        exit 1
    fi

    echo "Decrypting .env.age → .env using ${SSH_KEY}..."
    age -d -i "${SSH_KEY}" .env.age > .env

    if [ $? -eq 0 ]; then
        echo "✓ Successfully decrypted .env.age → .env"
    else
        echo "✗ Decryption failed. Make sure your SSH key is one of the registered recipients."
        rm -f .env
        exit 1
    fi
    exit 0
fi
# ─────────────────────────────────────────────────────────────────────────────

# ── Encrypt mode (default) ────────────────────────────────────────────────────
if [ ! -f .env ]; then
    echo "✗ .env not found in current directory"
    exit 1
fi

echo "Fetching SSH keys from GitHub and encrypting .env..."

(for user in "${RECIPIENTS[@]}"; do
    curl -s "https://github.com/${user}.keys"
done) | \
 age -R - .env > .env.age

if [ $? -eq 0 ]; then
    echo "✓ Successfully encrypted .env → .env.age"
else
    echo "✗ Encryption failed"
    rm -f .env.age
    exit 1
fi
```