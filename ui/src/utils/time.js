export const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp)

  const zeroPaddedHours =  date.getUTCHours() < 10 ? "0" + date.getUTCHours() : date.getUTCHours()
  const zeroPaddedMinutes = date.getUTCMinutes() < 10 ? "0" + date.getUTCMinutes() : date.getUTCMinutes()
  const zeroPaddedSeconds = date.getUTCSeconds() < 10 ? "0" + date.getUTCSeconds() : date.getUTCSeconds()

  return `${date.getUTCDate()}.${date.getUTCMonth()}.${date.getUTCFullYear()} ${zeroPaddedHours}:${zeroPaddedMinutes}:${zeroPaddedSeconds}`
}
