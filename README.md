# Homelab Transcriber

Go service for transcribing video clips using OpenAI Whisper API or local Whisper binary.

## A Part of My Future Go Series

In 2026, I'm writing each line by hand. The only AI allowed is for minor debugging and integrations within projects that require API use.

## Starting Point: August 17, 2026

This is where the transcriber journey begins. Everything after this date will build on this foundation.

### Why I Made This

I wanted a way to quickly find relationships between video clips. As I'm starting YouTube, I'm finding it very difficult to recall which clips say what and how clips that may be out of order relate to each other.

### Vision (Post 8/17)

A full application running on my homelab, connected to a database and dashboard. With an external hard drive, I'll store clips and save them over time. Depending on date and timestamps, I can easily edit clips together and pull from my phone, Canon camera, or straight from the clip database. This will be incredibly useful when newer clips relate to older ones.

## Timeline

**8/17/2026:** Initial transcriber MVP — accepts video upload, calls Whisper, returns JSON transcript
**Future:** Database integration, dashboard, multi-project scaling
