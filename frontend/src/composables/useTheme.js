import {computed, ref, watchEffect} from 'vue'

// 预定义主题配置
export const themes = {
    auto: {
        id: 'auto',
        name: '跟随系统',
        icon: '🌓',
        description: '根据系统设置自动切换明暗主题'
    },
    ocean: {
        id: 'ocean',
        name: '海洋蓝',
        icon: '🌊',
        description: '清新的海洋蓝色主题',
        colors: {
            primary: '#667eea',
            secondary: '#764ba2',
            accent: '#409eff',
            success: '#67c23a',
            warning: '#e6a23c',
            error: '#f56c6c'
        }
    },
    sunset: {
        id: 'sunset',
        name: '日落橙',
        icon: '🌅',
        description: '温暖的日落橙色主题',
        colors: {
            primary: '#ff6b6b',
            secondary: '#ffa726',
            accent: '#ff9800',
            success: '#4caf50',
            warning: '#ff9800',
            error: '#f44336'
        }
    },
    forest: {
        id: 'forest',
        name: '森林绿',
        icon: '🌲',
        description: '自然的森林绿色主题',
        colors: {
            primary: '#4caf50',
            secondary: '#81c784',
            accent: '#66bb6a',
            success: '#4caf50',
            warning: '#ff9800',
            error: '#f44336'
        }
    },
    purple: {
        id: 'purple',
        name: '紫罗兰',
        icon: '💜',
        description: '优雅的紫色主题',
        colors: {
            primary: '#9c27b0',
            secondary: '#ba68c8',
            accent: '#e91e63',
            success: '#4caf50',
            warning: '#ff9800',
            error: '#f44336'
        }
    },
    dark: {
        id: 'dark',
        name: '深色模式',
        icon: '🌙',
        description: '护眼的深色主题',
        colors: {
            primary: '#bb86fc',
            secondary: '#3700b3',
            accent: '#03dac6',
            success: '#4caf50',
            warning: '#ff9800',
            error: '#cf6679'
        }
    }
}

// 当前主题状态
const currentThemeId = ref('auto')
const systemPrefersDark = ref(false)

// 监听系统主题变化
const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
systemPrefersDark.value = mediaQuery.matches

mediaQuery.addEventListener('change', (e) => {
    systemPrefersDark.value = e.matches
})

// 计算当前实际主题
const currentTheme = computed(() => {
    if (currentThemeId.value === 'auto') {
        return systemPrefersDark.value ? themes.dark : themes.ocean
    }
    return themes[currentThemeId.value] || themes.ocean
})

// 主题管理 Composable
export function useTheme() {
    // 从本地存储加载主题设置
    const loadTheme = () => {
        try {
            const savedTheme = localStorage.getItem('mini-toolbox-theme')
            if (savedTheme && themes[savedTheme]) {
                currentThemeId.value = savedTheme
            }
        } catch (error) {
            console.warn('无法加载主题设置:', error)
        }
    }

    // 保存主题设置到本地存储
    const saveTheme = (themeId) => {
        try {
            localStorage.setItem('mini-toolbox-theme', themeId)
        } catch (error) {
            console.warn('无法保存主题设置:', error)
        }
    }

    // 设置主题
    const setTheme = (themeId) => {
        if (themes[themeId]) {
            currentThemeId.value = themeId
            saveTheme(themeId)
        }
    }

    // 应用主题到DOM
    const applyTheme = (theme) => {
        const root = document.documentElement

        if (theme.colors) {
            // 应用自定义颜色
            Object.entries(theme.colors).forEach(([key, value]) => {
                root.style.setProperty(`--theme-${key}`, value)
            })
        }

        // 设置主题标识
        root.setAttribute('data-theme', theme.id)

        // 如果是深色主题，添加dark类
        if (theme.id === 'dark' || (currentThemeId.value === 'auto' && systemPrefersDark.value)) {
            root.classList.add('dark-theme')
        } else {
            root.classList.remove('dark-theme')
        }
    }

    // 监听主题变化并应用
    watchEffect(() => {
        applyTheme(currentTheme.value)
    })

    // 初始化
    loadTheme()

    return {
        themes,
        currentThemeId: computed(() => currentThemeId.value),
        currentTheme,
        systemPrefersDark: computed(() => systemPrefersDark.value),
        setTheme,
        loadTheme
    }
}