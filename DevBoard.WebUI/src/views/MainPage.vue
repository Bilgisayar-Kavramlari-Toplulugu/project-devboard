<template>
  <div class="page">

    <!-- ── HERO ── -->
    <section class="hero">
      <div class="hero-badge">
        <span class="badge-dot"></span>
        20k+ Developers & Growing
      </div>

      <h1 class="hero-title">
        Discover Top
        <span class="gradient-text">Developer Talent</span>
        You Never Knew Existed
      </h1>

      <p class="hero-sub">
        DevBoard is the professional network built for engineers. Showcase your skills,
        explore others' work, and connect with the world's best developers.
      </p>

      <div class="hero-actions">
        <button class="cta-primary">Get Started Free</button>
        <button class="cta-ghost">
          <span class="play-icon">▷</span> See how it works
        </button>
      </div>

      <div class="hero-stats">
        <div class="stat">
          <span class="stat-num">20k+</span>
          <span class="stat-label">Developers</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat">
          <span class="stat-num">150+</span>
          <span class="stat-label">Countries</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat">
          <span class="stat-num">98%</span>
          <span class="stat-label">Satisfaction</span>
        </div>
      </div>
    </section>

    <!-- ── SECTION HEADER ── -->
    <section class="cards-section">
      <div class="section-header">
        <span class="section-badge">Featured Developers</span>
        <h2 class="section-title">Meet the Community</h2>
        <p class="section-sub">Top engineers sharing their craft, open to collaborate and connect.</p>
      </div>

      <!-- ── FILTER TABS ── -->
      <div class="filter-tabs">
        <button
          v-for="tab in tabs"
          :key="tab"
          :class="['tab-btn', { active: activeTab === tab }]"
          @click="activeTab = tab"
        >
          {{ tab }}
        </button>
      </div>

      <!-- ── DEVELOPER CARDS ── -->
      <div class="cards-grid">
        <div
          v-for="dev in filteredDevs"
          :key="dev.id"
          class="dev-card"
        >
          <!-- Card glow -->
          <div class="card-glow" :style="{ background: dev.glowColor }"></div>

          <!-- Top row -->
          <div class="card-top">
            <div class="avatar-wrap">
              <img :src="dev.avatar" :alt="dev.name" class="avatar" />
              <span class="online-dot" v-if="dev.online"></span>
            </div>
            <div class="card-meta">
              <h3 class="dev-name">{{ dev.name }}</h3>
              <p class="dev-role">{{ dev.role }}</p>
              <p class="dev-location">📍 {{ dev.location }}</p>
            </div>
            <button class="connect-btn">Connect</button>
          </div>

          <!-- Bio -->
          <p class="dev-bio">{{ dev.bio }}</p>

          <!-- Skills -->
          <div class="skills">
            <span
              v-for="skill in dev.skills"
              :key="skill"
              class="skill-tag"
            >{{ skill }}</span>
          </div>

          <!-- Stats -->
          <div class="card-stats">
            <div class="card-stat">
              <span class="cs-num">{{ dev.projects }}</span>
              <span class="cs-label">Projects</span>
            </div>
            <div class="card-stat">
              <span class="cs-num">{{ dev.followers }}</span>
              <span class="cs-label">Followers</span>
            </div>
            <div class="card-stat">
              <span class="cs-num">{{ dev.exp }}</span>
              <span class="cs-label">Years Exp</span>
            </div>
          </div>

          <!-- Links -->
          <div class="card-links">
            <a href="#" class="icon-link" title="GitHub">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"/>
              </svg>
            </a>
            <a href="#" class="icon-link" title="LinkedIn">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433a2.062 2.062 0 0 1-2.063-2.065 2.064 2.064 0 1 1 2.063 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
            </a>
            <a href="#" class="icon-link" title="Portfolio">
              <svg width="16" height="16" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8 8 3.589 8 8-3.589 8-8 8zm-.5-5h1v-5h-1v5zm0-6h1v-1h-1v1z"/>
              </svg>
            </a>
          </div>
        </div>
      </div>

      <div class="load-more">
        <button class="cta-ghost-lg">Load more developers</button>
      </div>
    </section>

    <!-- ── CTA BANNER ── -->
    <section class="cta-banner">
      <div class="cta-banner-inner">
        <h2>Ready to showcase your work?</h2>
        <p>Join 20,000+ developers already on DevBoard. Your profile is free forever.</p>
        <button class="cta-primary">Create your profile</button>
      </div>
    </section>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const tabs = ['All', 'Frontend', 'Backend', 'Full Stack', 'Mobile', 'DevOps']
const activeTab = ref('All')

const developers = ref([
  {
    id: 1,
    name: 'Alex Chen',
    role: 'Senior Frontend Engineer',
    location: 'San Francisco, CA',
    bio: 'Crafting pixel-perfect UIs with Vue & React. Open source contributor, design systems enthusiast.',
    skills: ['Vue', 'TypeScript', 'Figma', 'CSS'],
    category: 'Frontend',
    projects: 42,
    followers: '1.2k',
    exp: 7,
    online: true,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=alex&backgroundColor=b6e3f4',
    glowColor: 'radial-gradient(circle at top left, rgba(139,92,246,0.25) 0%, transparent 70%)',
  },
  {
    id: 2,
    name: 'Mia Torres',
    role: 'Full Stack Developer',
    location: 'Berlin, Germany',
    bio: 'Building scalable web apps from database to UI. Node.js, PostgreSQL and Vue.js daily driver.',
    skills: ['Node.js', 'Vue', 'PostgreSQL', 'Docker'],
    category: 'Full Stack',
    projects: 35,
    followers: '890',
    exp: 5,
    online: true,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=mia&backgroundColor=ffd5dc',
    glowColor: 'radial-gradient(circle at top left, rgba(236,72,153,0.2) 0%, transparent 70%)',
  },
  {
    id: 3,
    name: 'Kenji Tanaka',
    role: 'Backend Engineer',
    location: 'Tokyo, Japan',
    bio: 'Distributed systems, Rust & Go. I care about correctness, performance and clean APIs.',
    skills: ['Rust', 'Go', 'Kafka', 'Kubernetes'],
    category: 'Backend',
    projects: 28,
    followers: '2.1k',
    exp: 9,
    online: false,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=kenji&backgroundColor=c0aede',
    glowColor: 'radial-gradient(circle at top left, rgba(34,211,238,0.2) 0%, transparent 70%)',
  },
  {
    id: 4,
    name: 'Sara Yilmaz',
    role: 'Mobile Developer',
    location: 'Istanbul, Turkey',
    bio: 'React Native & Flutter. Building cross-platform apps that feel truly native. UI/UX focused.',
    skills: ['React Native', 'Flutter', 'Swift', 'Firebase'],
    category: 'Mobile',
    projects: 19,
    followers: '640',
    exp: 4,
    online: true,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=sara&backgroundColor=ffdfbf',
    glowColor: 'radial-gradient(circle at top left, rgba(251,146,60,0.2) 0%, transparent 70%)',
  },
  {
    id: 5,
    name: 'Lucas Müller',
    role: 'DevOps & Cloud Engineer',
    location: 'Amsterdam, Netherlands',
    bio: 'Infrastructure as code evangelist. AWS, Terraform, CI/CD pipelines and zero-downtime deploys.',
    skills: ['AWS', 'Terraform', 'Kubernetes', 'Python'],
    category: 'DevOps',
    projects: 31,
    followers: '1.5k',
    exp: 8,
    online: false,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=lucas&backgroundColor=d1d4f9',
    glowColor: 'radial-gradient(circle at top left, rgba(99,102,241,0.25) 0%, transparent 70%)',
  },
  {
    id: 6,
    name: 'Priya Sharma',
    role: 'Frontend / Design Engineer',
    location: 'Mumbai, India',
    bio: 'I live at the intersection of design and code. Accessibility, animations and design systems.',
    skills: ['React', 'Framer Motion', 'Storybook', 'WCAG'],
    category: 'Frontend',
    projects: 24,
    followers: '970',
    exp: 6,
    online: true,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=priya&backgroundColor=b6e3f4',
    glowColor: 'radial-gradient(circle at top left, rgba(168,85,247,0.25) 0%, transparent 70%)',
  },
])

const filteredDevs = computed(() =>
  activeTab.value === 'All'
    ? developers.value
    : developers.value.filter(d => d.category === activeTab.value)
)
</script>

<style scoped>
.page {
  display: flex;
  flex-direction: column;
}

/* ── HERO ── */
.hero {
  text-align: center;
  padding: 7rem 2rem 5rem;
  max-width: 860px;
  margin: 0 auto;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(168, 85, 247, 0.12);
  border: 1px solid rgba(168, 85, 247, 0.3);
  color: #c084fc;
  font-size: 0.8rem;
  font-weight: 600;
  padding: 0.35rem 1rem;
  border-radius: 999px;
  margin-bottom: 1.75rem;
  letter-spacing: 0.04em;
  text-transform: uppercase;
}

.badge-dot {
  width: 6px;
  height: 6px;
  background: #a855f7;
  border-radius: 50%;
  box-shadow: 0 0 6px #a855f7;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.4; }
}

.hero-title {
  font-size: clamp(2.2rem, 5vw, 3.6rem);
  font-weight: 800;
  line-height: 1.1;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 1.25rem;
}

.gradient-text {
  background: linear-gradient(135deg, #a855f7, #ec4899, #6366f1);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-sub {
  color: rgba(226, 232, 240, 0.6);
  font-size: 1.05rem;
  line-height: 1.7;
  margin-bottom: 2.25rem;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 3.5rem;
}

.cta-primary {
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  border: none;
  color: #fff;
  font-size: 0.95rem;
  font-weight: 600;
  padding: 0.75rem 1.75rem;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0 0 30px rgba(168, 85, 247, 0.4);
  transition: all 0.2s;
}

.cta-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 0 45px rgba(168, 85, 247, 0.6);
}

.cta-ghost {
  background: rgba(255,255,255,0.06);
  border: 1px solid rgba(255,255,255,0.1);
  color: rgba(226,232,240,0.85);
  font-size: 0.95rem;
  padding: 0.75rem 1.5rem;
  border-radius: 10px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.2s;
}

.cta-ghost:hover {
  background: rgba(255,255,255,0.1);
  color: #fff;
}

.play-icon {
  font-size: 0.8rem;
}

.hero-stats {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 2rem;
}

.stat {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.2rem;
}

.stat-num {
  font-size: 1.6rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
}

.stat-label {
  font-size: 0.75rem;
  color: rgba(226,232,240,0.45);
  text-transform: uppercase;
  letter-spacing: 0.06em;
}

.stat-divider {
  width: 1px;
  height: 36px;
  background: rgba(255,255,255,0.1);
}

/* ── CARDS SECTION ── */
.cards-section {
  padding: 4rem 2rem 5rem;
  max-width: 1280px;
  margin: 0 auto;
  width: 100%;
}

.section-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.section-badge {
  display: inline-block;
  background: rgba(168, 85, 247, 0.12);
  border: 1px solid rgba(168, 85, 247, 0.25);
  color: #c084fc;
  font-size: 0.75rem;
  font-weight: 600;
  padding: 0.3rem 0.9rem;
  border-radius: 999px;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  margin-bottom: 1rem;
}

.section-title {
  font-size: 2.2rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.75rem;
}

.section-sub {
  color: rgba(226,232,240,0.5);
  font-size: 0.95rem;
}

/* ── FILTER TABS ── */
.filter-tabs {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 2.5rem;
}

.tab-btn {
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.08);
  color: rgba(226,232,240,0.6);
  font-size: 0.85rem;
  padding: 0.45rem 1.1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.tab-btn:hover {
  color: #fff;
  background: rgba(255,255,255,0.08);
}

.tab-btn.active {
  background: linear-gradient(135deg, rgba(124,58,237,0.4), rgba(168,85,247,0.4));
  border-color: rgba(168,85,247,0.5);
  color: #e9d5ff;
  font-weight: 600;
}

/* ── CARDS GRID ── */
.cards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 1.5rem;
}

.dev-card {
  position: relative;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.07);
  border-radius: 16px;
  padding: 1.5rem;
  overflow: hidden;
  transition: border-color 0.25s, transform 0.25s;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.dev-card:hover {
  border-color: rgba(168, 85, 247, 0.3);
  transform: translateY(-3px);
}

.card-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  opacity: 0;
  transition: opacity 0.3s;
}

.dev-card:hover .card-glow {
  opacity: 1;
}

/* ── CARD TOP ── */
.card-top {
  display: flex;
  align-items: flex-start;
  gap: 0.85rem;
}

.avatar-wrap {
  position: relative;
  flex-shrink: 0;
}

.avatar {
  width: 52px;
  height: 52px;
  border-radius: 50%;
  border: 2px solid rgba(168,85,247,0.3);
  background: rgba(255,255,255,0.05);
}

.online-dot {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 10px;
  height: 10px;
  background: #22c55e;
  border-radius: 50%;
  border: 2px solid #0d0d1a;
}

.card-meta {
  flex: 1;
  min-width: 0;
}

.dev-name {
  font-size: 1rem;
  font-weight: 700;
  color: #fff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.dev-role {
  font-size: 0.8rem;
  color: #a855f7;
  font-weight: 500;
  margin-top: 0.15rem;
}

.dev-location {
  font-size: 0.75rem;
  color: rgba(226,232,240,0.4);
  margin-top: 0.2rem;
}

.connect-btn {
  flex-shrink: 0;
  background: rgba(168,85,247,0.12);
  border: 1px solid rgba(168,85,247,0.3);
  color: #c084fc;
  font-size: 0.78rem;
  font-weight: 600;
  padding: 0.35rem 0.85rem;
  border-radius: 7px;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.connect-btn:hover {
  background: rgba(168,85,247,0.25);
  color: #e9d5ff;
}

/* ── BIO ── */
.dev-bio {
  font-size: 0.83rem;
  color: rgba(226,232,240,0.55);
  line-height: 1.6;
}

/* ── SKILLS ── */
.skills {
  display: flex;
  flex-wrap: wrap;
  gap: 0.4rem;
}

.skill-tag {
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.08);
  color: rgba(226,232,240,0.7);
  font-size: 0.72rem;
  font-weight: 500;
  padding: 0.25rem 0.65rem;
  border-radius: 6px;
  letter-spacing: 0.02em;
}

/* ── CARD STATS ── */
.card-stats {
  display: flex;
  gap: 1rem;
  padding-top: 0.75rem;
  border-top: 1px solid rgba(255,255,255,0.05);
}

.card-stat {
  display: flex;
  flex-direction: column;
  gap: 0.1rem;
}

.cs-num {
  font-size: 1rem;
  font-weight: 700;
  color: #fff;
}

.cs-label {
  font-size: 0.7rem;
  color: rgba(226,232,240,0.4);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

/* ── CARD LINKS ── */
.card-links {
  display: flex;
  gap: 0.75rem;
  margin-top: auto;
}

.icon-link {
  color: rgba(226,232,240,0.35);
  transition: color 0.2s;
  display: flex;
}

.icon-link:hover {
  color: #a855f7;
}

/* ── LOAD MORE ── */
.load-more {
  text-align: center;
  margin-top: 3rem;
}

.cta-ghost-lg {
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.1);
  color: rgba(226,232,240,0.75);
  font-size: 0.9rem;
  padding: 0.75rem 2rem;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.2s;
}

.cta-ghost-lg:hover {
  background: rgba(255,255,255,0.08);
  color: #fff;
}

/* ── CTA BANNER ── */
.cta-banner {
  padding: 5rem 2rem;
}

.cta-banner-inner {
  max-width: 680px;
  margin: 0 auto;
  text-align: center;
  background: rgba(168,85,247,0.08);
  border: 1px solid rgba(168,85,247,0.2);
  border-radius: 20px;
  padding: 3.5rem 3rem;
  position: relative;
  overflow: hidden;
}

.cta-banner-inner::before {
  content: '';
  position: absolute;
  top: -40%;
  left: 50%;
  transform: translateX(-50%);
  width: 300px;
  height: 200px;
  background: radial-gradient(ellipse, rgba(168,85,247,0.25) 0%, transparent 70%);
  pointer-events: none;
}

.cta-banner-inner h2 {
  font-size: 2rem;
  font-weight: 800;
  color: #fff;
  letter-spacing: -0.03em;
  margin-bottom: 0.75rem;
}

.cta-banner-inner p {
  color: rgba(226,232,240,0.55);
  font-size: 0.95rem;
  margin-bottom: 1.75rem;
}
</style>
