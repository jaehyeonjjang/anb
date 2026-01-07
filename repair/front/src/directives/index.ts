import type { App } from 'vue';
import { setupInfiniteDirective } from './infinite';

export function setupGlobalDirectives(app: App) {
    setupInfiniteDirective(app)
}
