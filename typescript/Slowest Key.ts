function slowestKey(releaseTimes: number[], keysPressed: string): string {
  let maxDuration = releaseTimes[0];
  let maxKey = keysPressed[0];

  for (let i = 1; i < keysPressed.length; i++) {
    const duration = releaseTimes[i] - releaseTimes[i - 1];
    if (duration < maxDuration) {
      continue;
    }

    if (duration > maxDuration) {
      maxDuration = duration;
      maxKey = keysPressed[i];
      continue;
    }

    if (keysPressed[i] > maxKey) {
      maxKey = keysPressed[i];
    }
  }

  return maxKey;
}
